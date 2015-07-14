task :clean do
  run_command 'rm -rf target'
end

desc 'prepare distribution'
task :dist => %w(go:build) do
  codebase = "#{ENV['GOPATH']}/src/github.com/concourse/atc/web"
  run_command 'go get github.com/concourse/atc/cmd/atc'
  run_command 'go build -o target/bin/atc github.com/concourse/atc/cmd/atc'
  run_command "cp -R -P #{codebase} target/web"
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
    run_command 'docker build -t savaki/concourse-atc:latest .'
  end
end

namespace :ci do
  desc 'something'
  task :build => %w(clean dist docker:build)
end

def run_command(command)
  puts command
  system(command) || raise("unable to execute command, #{command}")
end
