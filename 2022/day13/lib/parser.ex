defmodule Parser do
  @spec lines(File.Stream) :: Stream.t()
  def lines(fstream) do
    fstream
    |> Stream.map(&String.trim/1)
  end

  @spec data(Stream.t()) :: Stream.t()
  def data(lines) do
    lines |> Enum.chunk_every(3) |> Enum.map(fn l -> parse_pair(l) end)
  end

  @spec parse_pair([String.t()]) :: {[], []}
  defp parse_pair(s) do
    [first, second] = s |> Enum.take(2)
    {parse_packet(first), parse_packet(second)}
  end

  @spec parse_packet(String.t()) :: []
  defp parse_packet(s) do
    Code.eval_string(s) |> elem(0)
  end
end
