defmodule Day14.CLI do
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

    IO.write("task #2 solution:")
    0 = Task2.solution(contents)
    IO.puts("<see images/>")
    {:ok, self()}
  end
end
