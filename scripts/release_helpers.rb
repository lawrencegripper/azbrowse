require 'colorize'
require 'English'
require 'fileutils'
require 'open3'

def print_header(title)
  # Used to create collapsing sections in github actions output
  github_actions_block = ['::endgroup::', "::group::#{title}"]
  title_block = [
    '',
    '------------------------------------------------------------------',
    "----> #{title}",
    '------------------------------------------------------------------'
  ]
  puts github_actions_block.join("\n").colorize(color: :black)
  title_text = title_block.map { |string| string.colorize(color: :white, background: :green) }.join("\n")
  puts title_text
end

def error_if_git_has_changes(git_instance, error_to_show_if_changes)
  git_changes = git_instance.status.changed
  return unless git_changes.count.positive?

  git_changes.keys { |key| puts "Changes in file #{key}" }
  exit_with_error error_to_show_if_changes
end

def exit_with_error(error)
  puts "Failed #{error}".colorize(color: :white, background: :red)
  raise error
end

# rubocop:disable Metrics/MethodLength
def execute_command(command)
  stdin, stdout_stderr, thread = Open3.popen2e(command.to_s)

  while stdout_stderr.closed? == false && line = stdout_stderr.gets
    puts(line)
  end
  thread_exit_status = thread.value
  stdin.close
  stdout_stderr.close

  puts thread_exit_status

  return if thread_exit_status.success?

  puts 'Failed'.colorize(color: :white, background: :red)
  raise "Command '#{command}' failed to execute"
end
