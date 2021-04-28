#!/usr/bin/env ruby

require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem 'git', '~> 1.8'
  gem 'colorize'
  gem 'docker-api'
end

require 'colorize'

def printHeader(title)
  background = :light_green
  puts "::endgroup::"
  puts "::group::#{title}"
  puts ""
  puts "------------------------------------------------------------------".colorize(:color => :light_blue, :background => background)
  puts "----> " + title.bold + "                                      ----".colorize(:color => :blue)
  puts "------------------------------------------------------------------".colorize(:color => :light_blue, :background => background)
  puts ""
end

def checkGitHasNoChanges(errorToShowIfChanges)
  if git_instance.status.changed.count > 0
    git_instance.status.changed.keys {|key| puts "Changes in file #{key}" }
    exitWithError errorToShowIfChanges
  end
end

def exitWithError(error)
  puts "Failed #{error}".colorize(:color => :white, :background => :red)
  puts ""
  exit(1)
end

# Move to the root of the repo
script_root = File.dirname(__FILE__)
repo_root = "#{script_root}/../"
Dir.chdir repo_root

printHeader("Setup - Checking things good to create a release")

required_envs = {
  'docker_username': 'DOCKER_USERNAME',
  'docker_password': 'DOCKER_PASSWORD',
  'github_token': 'GITHUB_TOKEN',
  'snapcraft_login': 'SNAPCRAFT_LOGIN',
  'is_ci': 'IS_CI',
  'branch': 'BRANCH',
  'build_number': 'BUILD_NUMBER',
}
required_envs.each do |var_name, env_name|
  exitWithError "Missing required ENV #{env_name}" if ENV[env_name] == "" or ENV[env_name] == nil
  instance_variable_set("@#{var_name}", ENV[env_name])
end

printHeader("Configuration")
puts "Is running in CI? #{@is_ci}"
puts "Branch: #{@branch}"

# By default don't publish build output
publish_build_output = false

if @is_ci and branch == "/refs/heads/main"
  publish_build_output = true
  puts "Login to docker".colorize(:blue)
  Docker.authenticate!('username' => ENV['DOCKER_USER'], 'password' => ENV['DOCKER_PASSWORD'])
  Docker.authenticate!('serveraddress' => 'ghcr.io', 'username' => ENV['DOCKER_USER'], 'password' => ENV['GITHUB_TOKEN'])
  
  puts "Login to snapcraft".colorize(:blue)
  puts ` 
    echo $SNAPCRAFT_LOGIN | base64 -d > snap.login
    snapcraft login --with snap.login
    # cleanup login file
    rm snap.login
  `
else
  puts "Skipping publish as either not CI or branch != main"
end

printHeader('Git - Create tag for release')
tag = "v2.0.#{@build_number}"
puts "Tag: #{tag}"
git_instance = Git.open(repo_root)
git_instance.add_tag(tag)

printHeader('Build, lint and codegen')
puts `make ci`
checkGitHasNoChanges('Codegen caused changes to files. Run "make swagger-codegen" and commit the results to resolve this issue')

printHeader('Generate docs')
`make docs-update`
checkGitHasNoChanges('Docs generation caused git changes. Run "make docs-update" and commit the results to resolve this issue.')

printHeader('Run integration tests')
`./scripts/ci_integration_tests.sh`

printHeader('Run goreleaser')
if publish_build_output 
  puts `goreleaser --skip-publish --rm-dist`
else
  puts `goreleaser`
  docker push "$DEV_CONTAINER_TAG"
end

# Push up built output for the devcontainer if we're on main
devcontainer_images = [
  'ghcr.io/lawrencegripper/azbrowse/devcontainer:latest',
  'ghcr.io/lawrencegripper/azbrowse/snapbase:latest'
]
devcontainer_images.each do |image_name| 
  if publish_build_output and Docker::Image.exist?(image_name)
    printHeader("Push devcontainer image #{image_name}")
    Docker::Image.get(image_name).push()
  end
end

