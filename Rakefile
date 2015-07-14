MAJOR = 0
MINOR = 1

task :clean do
  run_command 'rm -rf target'
end

def version
  # if ENV['CIRCLE_BUILD_NUM'].nil? or ENV['CIRCLE_SHA1'].nil?
    'latest'
  # else
  #   "#{MAJOR}.#{MINOR}.#{ENV['CIRCLE_BUILD_NUM']}.#{ENV['CIRCLE_SHA1'][0..5]}"
  # end
end

desc 'prepare distribution'
task :dist => %w(go:build) do
  run_command 'go get github.com/concourse/atc/cmd/atc'
  run_command 'go build -o target/bin/atc github.com/concourse/atc/cmd/atc'
  run_command 'chmod +x target/atc'

  run_command 'git clone https://github.com/concourse/atc'
  run_command 'cp -R -P atc/web target/web'
end

namespace :go do
  desc 'builds the atcd binary'
  task :build do
    run_command '(cd cmd/atcd; go get ; go build -o ../../target/bin/atcd)'
  end
end

namespace :docker do
  desc 'builds the docker image'
  task :build do
    run_command "docker build -t savaki/concourse-atc:#{version} ."
  end

  desc 'pushes the docker image'
  task :push do
    run_command "docker push savaki/concourse-atc:#{version}"
  end
end

namespace :ci do
  desc 'builds docker image in ci environments'
  task :build => %w(clean dist docker:build)

  desc 'pushes image to docker hub'
  task :push => %w(docker:push)
end

def run_command(command)
  puts command
  system(command) || raise("unable to execute command, #{command}")
end
