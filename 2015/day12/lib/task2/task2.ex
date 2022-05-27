defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(line) do
    solveSimpleObjects(line, 0)
  end

  def solveSimpleObjects(line, sum) do
    case Regex.match?(~r/^\{[^\{\}\]\[]*\}$/, line) do
      true ->
        case String.contains?(line, "red") do
          true ->
            0

          false ->
            sum + Task1.solution(line)
        end

      false ->
        simpleObjs = Regex.scan(~r/\{[^\{\}\]\[]+\}/, line)

        case simpleObjs do
          [] ->
            solveSimpleArrays(line, sum)

          matches ->
            tempSum =
              Enum.filter(matches, fn [x] -> not String.contains?(x, "red") end)
              |> Enum.map(fn [x] -> Task1.solution(x) end)
              |> Enum.sum()

            solveSimpleObjects(Regex.replace(~r/\{[^\{\}\]\[]+\}/, line, "0"), sum + tempSum)
        end
    end
  end

  def solveSimpleArrays(line, sum) do
    case Regex.match?(~r/^\[[^\{\}\]\[]*\]$/, line) do
      true ->
        sum + Task1.solution(line)

      false ->
        simpleArrs = Regex.scan(~r/\[[^\{\}\]\[]+\]/, line)

        case simpleArrs do
          [] ->
            solveSimpleObjects(line, sum)

          matches ->
            tempSum =
              Enum.map(matches, fn [x] -> Task1.solution(x) end)
              |> Enum.sum()

            solveSimpleObjects(Regex.replace(~r/\[[^\{\}\]\[]+\]/, line, "0"), sum + tempSum)
        end
    end
  end

  # array that doesn't contain array ~r/\[[^\[\]]+\]/
  # object that doesn't contain object ~r/\{[^\{\}]+\}/
  # object that doesn't contain objects and arrays ~r/\{[^\{\}\]\[]+\}/

  # line = ~s({"d":{"f":["red"]},"g":{"z":"asd"},"g":{"z":"asd"},"e":[1,2,3,4][,"f":5]]})
  # Regex.run(~r/\{[^\{\}\]\[]+\}/, line)
end
