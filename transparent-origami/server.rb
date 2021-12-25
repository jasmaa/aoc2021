require 'sinatra'
require './scenario'

set :bind, "0.0.0.0"
set :port, 8080

post '/first-fold-dots' do
  payload = request.body.read
  begin
    scenario = Scenario.parse(payload)
    if scenario.steps_left?
      scenario.step()
      content_type :json
      {
        dots: scenario.points.length
      }.to_json
    end
  rescue
    content_type :json
    status 400
    {
      error: "invalid input"
    }.to_json
  end
end

post '/code' do
  payload = request.body.read
  begin
    scenario = Scenario.parse(payload)
    while scenario.steps_left?
      scenario.step()
    end
    scenario.to_ascii
  rescue
    content_type :json
    status 400
    {
      error: "invalid input"
    }.to_json
  end
end