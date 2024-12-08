defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> parse()
    |> Enum.flat_map(fn {res, args} -> match(args, res) end)
    |> Enum.uniq()
    |> Enum.sum()
  end

  @spec match(any(), [number()], any()) :: [...]
  def match(prev, [], res) do
    if prev == res do
      [res]
    else
      [0]
    end
  end

  def match(prev, [arg | []], res) do
    if Enum.any?(
         [
           prev * arg == res,
           prev + arg == res
         ],
         & &1
       ) do
      [res]
    else
      [0]
    end
  end

  def match(prev, [arg | tail], res) do
    match(prev * arg, tail, res) ++ match(prev + arg, tail, res)
  end

  def match([arg | [second_arg | tail]], res) do
    match(arg * second_arg, tail, res) ++ match(arg + second_arg, tail, res)
  end

  defp parse(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(fn row ->
      row
      |> String.split(":", trim: true)
    end)
    |> Enum.map(fn [res, args] ->
      {
        String.to_integer(res),
        args |> String.split(" ", trim: true) |> Enum.map(fn num -> String.to_integer(num) end)
      }
    end)
  end
end
