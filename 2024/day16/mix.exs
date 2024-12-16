defmodule Dayt.MixProject do
  use Mix.Project

  def project do
    [
      app: :day16,
      version: "0.1.0",
      elixir: "~> 1.17",
      start_permanent: Mix.env() == :prod,
      deps: deps(),
      escript: escript()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      mod: {Day16.CLI, []},
      extra_applications: [:logger]
    ]
  end

  defp escript do
    [main_module: Day16.CLI]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:libgraph, "~> 0.16.0"}
    ]
  end
end
