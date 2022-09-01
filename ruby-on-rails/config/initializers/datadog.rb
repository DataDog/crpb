require 'ddtrace'

Datadog.configure do |c|
  c.tracing.instrument :rails
end