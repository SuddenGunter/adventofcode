defmodule Task2Runner do
  def main do
    case File.read("data.txt") do
      {:ok, contents} -> process(contents)
      {:error, reason} -> "failed to read file 'data.txt': #{reason}"
    end
    |> IO.puts()
  end

  def process(contents) do
    Task2.solution(
      contents
      |> String.split("\n", trim: true)
      |> Enum.map(fn line ->
        [first, second] = String.split(String.slice(line, 0..-1), " ", trim: true)
        {String.to_atom(first), String.to_integer(second)}
      end)
    )
  end
end
