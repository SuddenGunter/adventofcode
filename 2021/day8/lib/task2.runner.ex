defmodule Task1Runner do
  def main do
    case File.read("data.txt") do
      {:ok, contents} -> contents |> process
      {:error, reason} -> "failed to read file 'data.txt': #{reason}"
    end
    |> IO.puts()
  end

  def process(contents) do
    contents
    |> String.split("\n", trim: true)
    |> Task2.solution()
  end
end
