defmodule Task2Runner do
  def main do
    case File.read("data.txt") do
      {:ok, contents} -> process(contents)
      {:error, reason} -> "failed to read file 'data.txt': #{reason}"
    end
    |> IO.puts()
  end

  def process(contents) do
    contents
    |> String.trim_trailing()
    |> String.split(",", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Task2.solution()
  end
end
