defmodule Day15.CLI do
  use Application

  def start(_type, _args) do
    case Input.read("data.txt") do
      {:ok, contents} -> solve(contents)
      {:error, reason} -> IO.puts(reason)
    end
  end

  defp solve(contents) do
    IO.write("task #1 solution: ")
    Task1.solution(contents) |> IO.puts()

    IO.write("task #2 solution: ")
    Task2.solution(contents) |> IO.puts()

    {:ok, self()}
  end
end
