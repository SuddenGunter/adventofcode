defmodule Day10.CLI do
  def main(_args) do
    File.stream!("data.txt") |> Parser.data() |> solve()
    :ok
  end

  defp solve(contents) do
    IO.write("task #1 solution: ")
    Task1.solution(contents) |> IO.puts()

    IO.puts("task #2 solution: ")
    Task2.solution(contents)
  end
end
