Rails.application.routes.draw do
  namespace 'api' do
    namespace 'v1' do
      get 'players', to: 'players#index'
end
