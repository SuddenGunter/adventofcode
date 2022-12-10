defmodule Parser do
  @spec lines(File.Stream) :: Stream.t()
  def lines(fstream) do
    fstream
    |> Stream.map(&String.trim/1)
  end

  @spec lines(Stream.t()) :: Stream.t()
  def data(lines) do
    lines |> Stream.map(&read_command/1)
  end

  defp read_command(l) do
    case String.split(l, " ") do
      [_, arg] -> {:addx, as_int(arg)}
      [_] -> {:noop}
    end
  end

  defp as_int(str) do
    {n, _} = Integer.parse(str)
    n
  end
end
