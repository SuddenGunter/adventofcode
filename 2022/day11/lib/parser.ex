defmodule Parser do
  @spec lines(File.Stream) :: Stream.t()
  def lines(fstream) do
    fstream
    |> Stream.map(&String.trim/1)
  end

  @spec data(Stream.t()) :: Stream.t()
  def data(lines) do
    lines |> Enum.chunk_every(7) |> Enum.map(fn l -> init_monkey(l) end)
  end

  @spec init_monkey([String.t()]) :: Monkey.t()
  defp init_monkey(lines) do
    id = first_integer(Enum.at(lines, 0))

    items =
      Regex.scan(~r/\d+/, Enum.at(lines, 1))
      |> List.flatten()
      |> Enum.map(fn x ->
        {n, _} = Integer.parse(x)
        n
      end)

    operation = parse_op(Enum.at(lines, 2))
    test = parse_test(Enum.at(lines, 3), Enum.at(lines, 4), Enum.at(lines, 5))

    Monkey.new(id, items, operation, test)
  end

  @spec parse_op(String.t()) :: (integer() -> integer())
  defp parse_op(s) do
    [_, op, val] = Regex.run(~r/new = old (.) (.*)/, s)

    if val == "old" do
      case op do
        "+" -> fn x -> x + x end
        "*" -> fn x -> x * x end
        "-" -> fn x -> x - x end
        "/" -> fn x -> x / x end
      end
    else
      {num, _} = Integer.parse(val)

      case op do
        "+" -> fn x -> x + num end
        "*" -> fn x -> x * num end
        "-" -> fn x -> x - num end
        "/" -> fn x -> x / num end
      end
    end
  end

  @spec parse_test(String.t(), String.t(), String.t()) :: (integer() -> integer())
  defp parse_test(s, success, fail) do
    [_, val] = Regex.run(~r/divisible by (\d+)/, s)
    {n, _} = Integer.parse(val)

    if_true = first_integer(success)
    if_fail = first_integer(fail)

    fn x ->
      if rem(x, n) == 0 do
        if_true
      else
        if_fail
      end
    end
  end

  @spec first_integer(String.t()) :: integer()
  defp first_integer(s) do
    Regex.run(~r/\d+/, s) |> Enum.at(0) |> Integer.parse() |> elem(0)
  end
end
