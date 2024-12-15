defmodule Dayt.MixProject do
  use Mix.Project

  def project do
    [
      app: :day14,
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
      mod: {Day14.CLI, []},
      extra_applications: [:logger]
    ]
  end

  defp escript do
    [main_module: Day14.CLI]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    []
  end
end
