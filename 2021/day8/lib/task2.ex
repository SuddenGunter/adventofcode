use Bitwise

defmodule Task2 do
  def solution(inputs) do
    Enum.filter(inputs, fn x -> String.length(x) > 0 end)
    |> Enum.map(&processLine/1)
    |> Enum.sum()
  end

  def processLine(line) do
    split = String.split(line, " | ")
    dict = generateDictionary(Enum.at(split, 0))

    mapResult(dict, Enum.at(split, 1))
  end

  def generateDictionary(line) do
    knownByCounting = countLettersFrequency(line)
    inferredFromOne = oneHeuristic(line, Map.keys(knownByCounting))
    inferredFromOneAndSeven = oneSevenHeuristic(line)

    inferredFromFour =
      fourHeuristic(
        line,
        Map.keys(
          Map.merge(knownByCounting, inferredFromOne)
          |> Map.merge(inferredFromOneAndSeven)
        )
      )

    eight =
      eightHeuristic(
        line,
        Map.keys(
          Map.merge(knownByCounting, inferredFromOne)
          |> Map.merge(inferredFromOneAndSeven)
          |> Map.merge(inferredFromFour)
        )
      )

    Map.merge(knownByCounting, inferredFromOne)
    |> Map.merge(inferredFromOneAndSeven)
    |> Map.merge(inferredFromFour)
    |> Map.merge(eight)

    # %{
    #   "a" => "a",
    #   "b" => "b",
    #   "c" => "c",
    #   "d" => "d",
    #   "e" => "e",
    #   "f" => "f"
    # }
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

  def oneHeuristic(line, knownLetters) do
    numberOne =
      line
      |> String.split(" ")
      |> Enum.sort_by(fn x -> String.length(x) end)
      |> Enum.at(0)
      |> String.graphemes()

    %{((numberOne -- knownLetters) |> Enum.at(0)) => "c"}
  end

  def oneSevenHeuristic(line) do
    letters =
      line
      |> String.split(" ")
      |> Enum.sort_by(fn x -> String.length(x) end)
      |> Enum.take(2)
      |> Enum.map(&String.graphemes/1)

    %{((Enum.at(letters, 1) -- Enum.at(letters, 0)) |> Enum.at(0)) => "a"}
  end

  def fourHeuristic(line, knownLetters) do
    numberOne =
      line
      |> String.split(" ")
      |> Enum.sort_by(fn x -> String.length(x) end)
      |> Enum.at(2)
      |> String.graphemes()

    %{((numberOne -- knownLetters) |> Enum.at(0)) => "d"}
  end

  def eightHeuristic(line, knownLetters) do
    numberOne =
      line
      |> String.split(" ")
      |> Enum.sort_by(fn x -> String.length(x) end)
      |> List.last()
      |> String.graphemes()

    %{((numberOne -- knownLetters) |> Enum.at(0)) => "g"}
  end

  def mapResult(dict, right) do
    nums = String.split(right, " ")

    Enum.map(nums, fn x ->
      mapDigit(x, dict)
    end)
    |> IO.inspect()

    8394
  end

  def mapDigit(digit, dict) do
    bitShift = %{
      "a" => 0,
      "b" => 1,
      "c" => 2,
      "d" => 3,
      "e" => 4,
      "f" => 5,
      "g" => 6
    }

    binaryEncoded = %{
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

    digit |> IO.inspect(label: "digit")
    dict |> IO.inspect(label: "dict")

    mapped =
      String.graphemes(digit)
      |> Enum.map(fn x -> dict[x] end)

    mapped |> IO.inspect(label: "mapped")

    key =
      mapped
      |> Enum.reduce(
        0,
        fn x, acc ->
          acc ||| 1 <<< bitShift[x]
        end
      )

    key |> IO.inspect(label: "key")
    binaryEncoded[<<key::8>>] |> IO.inspect(label: "binaryEncoded[key]")
  end
end
