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

devcontainer_images = [
    'ghcr.io/lawrencegripper/azbrowse/devcontainer:latest',
    'ghcr.io/lawrencegripper/azbrowse/snapbase:latest'
  ]
  devcontainer_images.each do |image_name|
    if Docker::Image.exist?(image_name)
      print_header("Push devcontainer image #{image_name}")
      puts execute_command("docker push #{image_name}")
    end
  end

# git_instance = Git.open('.')
# last_release_tag = git_instance.describe("HEAD", {:tags => true, :abbrev => '0'})
# puts "Last release tag was #{last_release_tag}"

# changes_since_last_release = git_instance.gtree(last_release_tag).diff('HEAD').map(&:path)
# puts 'Changed files since last release:'
# puts changes_since_last_release
