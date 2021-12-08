use Bitwise

defmodule Task2 do
  def solution(inputs) do
    Enum.filter(inputs, fn x -> String.length(x) > 0 end)
    |> Enum.map(&processLine/1)
    |> Enum.sum()
  end

  def processLine(line) do
    [left, right] = String.split(line, " | ")
    generateDictionary(left) |> mapResult(right)
  end

  def generateDictionary(line) do
    knownByCounting = countLettersFrequency(line)

    words =
      line
      |> String.split(" ")
      |> Enum.sort_by(fn x -> String.length(x) end)

    # can be found by position
    %{
      0 => "c",
      1 => "a",
      2 => "d",
      # last
      (length(words) - 1) => "g"
    }
    |> Enum.reduce(
      knownByCounting,
      fn {k, v}, acc ->
        graphemes = words |> Enum.at(k) |> String.graphemes()

        Map.put_new(
          acc,
          (graphemes -- Map.keys(acc)) |> Enum.at(0),
          v
        )
      end
    )
  end

  def countLettersFrequency(line) do
    frq = line |> String.graphemes() |> Enum.filter(fn x -> x != " " end) |> Enum.frequencies()

    Enum.reduce(
      frq,
      %{},
      fn {k, v}, acc ->
        case v do
          4 -> Map.put_new(acc, k, "e")
          9 -> Map.put_new(acc, k, "f")
          6 -> Map.put_new(acc, k, "b")
          _ -> acc
        end
      end
    )
  end

  def mapResult(dict, right) do
    String.split(right, " ")
    |> Enum.map(fn x ->
      mapDigit(x, dict)
    end)
    |> Enum.reverse()
    |> Enum.with_index()
    |> Enum.reduce(0, fn {number, index}, acc ->
      acc + number * :math.pow(10, index)
    end)
    |> trunc
  end

  @bitShift %{
    "a" => 0,
    "b" => 1,
    "c" => 2,
    "d" => 3,
    "e" => 4,
    "f" => 5,
    "g" => 6
  }

  @binaryEncoded %{
    <<0::1, 1::1, 1::1, 1::1, 1::1, 1::1, 1::1, 1::1>> => 8,
    <<0::1, 1::1, 1::1, 1::1, 0::1, 1::1, 1::1, 1::1>> => 0,
    <<0::1, 1::1, 1::1, 1::1, 1::1, 0::1, 1::1, 1::1>> => 6,
    <<0::1, 1::1, 1::1, 0::1, 1::1, 1::1, 1::1, 1::1>> => 9,
    <<0::1, 1::1, 1::1, 0::1, 1::1, 0::1, 1::1, 1::1>> => 5,
    <<0::1, 1::1, 1::1, 0::1, 1::1, 1::1, 0::1, 1::1>> => 3,
    <<0::1, 1::1, 0::1, 1::1, 1::1, 1::1, 0::1, 1::1>> => 2,
    <<0::1, 0::1, 1::1, 0::1, 0::1, 1::1, 0::1, 1::1>> => 7,
    <<0::1, 0::1, 1::1, 0::1, 0::1, 1::1, 0::1, 0::1>> => 1,
    <<0::1, 0::1, 1::1, 0::1, 1::1, 1::1, 1::1, 0::1>> => 4
  }

  def mapDigit(digit, dict) do
    key =
      String.graphemes(digit)
      |> Enum.map(fn x -> dict[x] end)
      |> Enum.reduce(
        0,
        fn x, acc ->
          acc ||| 1 <<< @bitShift[x]
        end
      )

    @binaryEncoded[<<key::8>>]
  end
end
