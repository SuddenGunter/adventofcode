defmodule Parser do
  def lines(contents) do
    contents
    |> String.downcase()
    |> String.split("\n", trim: true)
  end

  @spec lines(String.t()) :: [%Reindeer{}]
  def data(contents) do
    contents
    |> lines
    |> Enum.map(fn x -> find_captures(x) end)
    |> Enum.filter(fn x -> x != nil end)
  end

  defp find_captures(line) do
    case Regex.scan(~r/^[^0-9]+(\d+)[^0-9]+(\d+)[^0-9]+(\d+)[^0-9]+$/, line) do
      [] ->
        nil

      [[_, speed, runTime, restTime]] ->
        %Reindeer{
          speed: String.to_integer(speed),
          runTime: String.to_integer(runTime),
          restTime: String.to_integer(restTime)
        }
    end
  end
end
