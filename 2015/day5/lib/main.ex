defmodule Day5.CLI do
  def main(_args) do
    case Input.read("data.txt", &Parser.lines/1) do
      {:ok, contents} -> solve(contents)
      {:error, reason} -> IO.puts(reason)
    end
  end

  defp solve(contents) do
    IO.write("task #1 solution: ")
    Task1.solution(contents) |> IO.puts()

    IO.write("task #2 solution: ")
    Task2.solution(contents) |> IO.puts()
  end
end
