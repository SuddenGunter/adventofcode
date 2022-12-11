defmodule Day11.CLI do
  def main(_args) do
    File.stream!("demo.data.txt") |> Parser.data() |> solve()
    :ok
  end

  defp solve(contents) do
    IO.write("task #1 solution: ")
    Task1.solution(contents) |> IO.puts()

    IO.write("task #2 solution: ")
    Task1.solution(contents) |> IO.puts()
  end
end
