defmodule Day1.CLI do
  def main(_args) do
    case Input.read("data.txt", &Parser.firstLine/1) do
      {:ok, contents} -> solve(contents)
      {:error, reason} -> IO.puts(reason)
    end
  end

  defp solve(contents) do
    Task1.solution(contents) |> IO.puts()
  end
end
