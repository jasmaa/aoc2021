require 'sinatra'
require './scenario'
require './scenario_optimized'

set :bind, "0.0.0.0"
set :port, 8080

post '/10' do
  payload = request.body.read
  begin
    scenario = Scenario.parse(payload)
    (1..10).each do |_|
      scenario.step()
    end
    content_type :json
    {
      difference: scenario.commonality_difference
    }.to_json
  rescue
    content_type :json
    {
      error: "invalid input"
    }.to_json
  end
end

post '/40' do
  payload = request.body.read
  begin
    scenario = ScenarioOptimized.parse(payload)
    (1..40).each do |_|
      scenario.step()
    end
    content_type :json
    {
      difference: scenario.commonality_difference
    }.to_json
  rescue
    content_type :json
    {
      error: "invalid input"
    }.to_json
  end
end