#!/usr/bin/env ruby

require 'bundler/inline'

gemfile do
  source 'https://rubygems.org'
  gem 'git', '~> 1.8'
  gem 'simple_cli'
  gem 'colorize'
  gem 'docker-api'
end

require 'simple_cli'
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

def exitWithError(error)
  puts "Failed #{error}".colorize(:color => :white, :background => :red)
  exit(1)
end

printHeader("Setup - Checking things good to create a release")

required_envs = {
  'docker_username' => 'DOCKER_USERNAME',
  'docker_password' => 'DOCKER_PASSWORD',
  # 'GITHUB_TOKEN', 'IS_CI', 'BRANCH'
}
required_envs.each do |var_name, env_name|
  exitWithError "Missing required ENV #{env_name}" if not ENV[env_name]
  instance_variable_set("@#{var_name}", ENV[env_name])
end

is_release = ENV["BUILD_NUMBER"]
if is_release == nil
  puts "Env 'BUILD_NUMBER must be set"
end

is_ci_pipeline = ENV["IS_CI"]
branch = ENV["BRANCH"]

if is_ci_pipeline and branch == "/refs/heads/main"
  ENV['PUBLISH'] = "true"
  should_publish = true
  
  puts "Login to docker".colorize(:blue)
  Docker.authenticate!('username' => ENV['DOCKER_USER'], 'password' => ENV['DOCKER_PASSWWORD'])
  Docker.authenticate!('serveraddress' => 'ghcr.io', 'username' => ENV['DOCKER_USER'], 'password' => ENV['GITHUB_TOKEN'])
  
  puts "Login to snapcraft".colorize(:blue)
  puts ` 
    echo $SNAPCRAFT_LOGIN | base64 -d > snap.login
    snapcraft login --with snap.login
    # cleanup login file
    rm snap.login
  `
end

