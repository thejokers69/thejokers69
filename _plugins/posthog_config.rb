# Jekyll plugin to load PostHog API key from environment variable
Jekyll::Hooks.register :site, :after_init do |site|
  if ENV['POSTHOG_API_KEY']
    site.config['posthog']['api_key'] = ENV['POSTHOG_API_KEY']
  end
end

