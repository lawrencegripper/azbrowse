#!/usr/bin/env ruby

require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem 'git', '=1.8.1'
  gem 'colorize', '=0.8.1'
  gem 'docker-api', '=2.1.0'
end

require 'colorize'
require 'English'
require 'fileutils'
require 'open3'
require_relative 'release_helpers'

begin
  # Ensure output is written to stdout without buffer
  $stdout.sync = true

  # Move to the root of the repo
  script_root = File.dirname(__FILE__)
  repo_root = "#{script_root}/../"
  Dir.chdir repo_root

  print_header('Setup - Checking things good to create a release')

  required_envs = {
    'is_ci': 'IS_CI',
    'branch': 'BRANCH',
    'build_number': 'BUILD_NUMBER'
  }

  required_envs_release = {
    'docker_username': 'DOCKER_USERNAME',
    'docker_password': 'DOCKER_PASSWORD',
    'github_token': 'GITHUB_TOKEN',
    'snapcraft_login': 'SNAPCRAFT_LOGIN'
  }

  required_envs = required_envs.merge(required_envs_release) if @branch == 'refs/heads/main'

  required_envs.each do |var_name, env_name|
    exit_with_error "Missing required ENV #{env_name}" if (ENV[env_name] == '') || ENV[env_name].nil?
    instance_variable_set("@#{var_name}", ENV[env_name])
  end

  # Get golang version, injected into build binaries for debugging
  go_version = `go version`
  ENV['GOVERSION'] = go_version

  print_header('Configuration')
  puts "Is running in CI? #{@is_ci}"
  puts "Branch: #{@branch}"
  puts "Go version: #{go_version}"

  # By default don't publish build output
  publish_build_output = false

  if @is_ci == 'true' && @branch == 'refs/heads/main'
    publish_build_output = true
    puts 'Login docker cli'.colorize(:blue)
    execute_command('./scripts/docker_login.sh')

    puts 'Login to snapcraft'.colorize(:blue)
    execute_command "
    echo $SNAPCRAFT_LOGIN | base64 -d > snap.login
    snapcraft login --with snap.login
    # cleanup login file
    rm snap.login
  "
  else
    puts 'Skipping publish as either not CI or branch != main'
  end

  puts "Is release to be published? #{publish_build_output}"

  git_instance = Git.open(repo_root)

  print_header('Build, lint and codegen')
  execute_command 'make ci'
  error_if_git_has_changes(git_instance,
                           'Codegen caused changes to files. Run "make swagger-codegen" and commit the results to resolve this issue')

  print_header('Generate docs')
  execute_command 'make docs-update'
  error_if_git_has_changes(git_instance,
                           'Docs generation caused git changes. Run "make docs-update" and commit the results to resolve this issue.')

  # Ensure race condition in snapcraft isn't expose
  # https://github.com/goreleaser/goreleaser/issues/1715
  FileUtils.mkdir_p("/#{ENV['HOME']}/.cache/snapcraft/download")
  FileUtils.mkdir_p("/#{ENV['HOME']}/.cache/snapcraft/stage-packages")
  if publish_build_output
    print_header('Git - Create tag for release')
    tag = "v2.1.#{@build_number}"
    puts "Tag: #{tag}"
    git_instance.add_tag(tag)
    puts "Push tag: #{tag}"
    git_instance.push('origin', 'main', { tags: true })

    print_header('Run goreleaser: Publish')
    execute_command 'goreleaser'
  else
    print_header('Run goreleaser: Dry run')
    execute_command 'goreleaser --skip-publish --rm-dist'
  end

  # Push up built output for the devcontainer if we're on main
  devcontainer_images = [
    'ghcr.io/lawrencegripper/azbrowse/devcontainer:latest',
    'ghcr.io/lawrencegripper/azbrowse/snapbase:latest'
  ]
  devcontainer_images.each do |image_name|
    if publish_build_output && Docker::Image.exist?(image_name)
      print_header("Push devcontainer image #{image_name}")
      Docker::Image.get(image_name).push
    end
  end
rescue StandardError => e
  puts ''
  puts 'Failure details:'
  puts e.message
  puts e.backtrace.inspect
  exit(1)
end
