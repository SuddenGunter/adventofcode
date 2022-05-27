defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(line) do
    solveSimpleObjects(line)
  end

  def solveSimpleObjects(line) do
    case Regex.match?(~r/^\{[^\{\}\]\[]*\}$/, line) do
      true ->
        # IO.inspect(line, label: "finish_simple_object")

        case String.contains?(line, "red") do
          true ->
            0

          false ->
            Task1.solution(line)
        end

      false ->
        simpleObjs = Regex.scan(~r/\{[^\{\}\]\[]+\}/, line)

        case simpleObjs do
          [] ->
            solveSimpleArrays(line)

          matches ->
            tempSums =
              Enum.map(matches, fn [x] ->
                case String.contains?(x, "red") do
                  true ->
                    0

                  false ->
                    Task1.solution(x)
                end
              end)

            replaced = Regex.replace(~r/\{[^\{\}\]\[]+\}/, line, "_replaced_")

            reduced =
              Enum.reduce(tempSums, replaced, fn x, acc ->
                Regex.replace(~r/_replaced_/, acc, Integer.to_string(x), global: false)
              end)

            # IO.write("line")
            # IO.puts(line)
            # IO.write("reduced")
            # IO.puts(reduced)

            solveSimpleObjects(reduced)
        end
    end
  end

  def solveSimpleArrays(line) do
    case Regex.match?(~r/^\[[^\{\}\]\[]*\]$/, line) do
      true ->
        line
        # |> IO.inspect(label: "finish_simple_array")
        |> Task1.solution()

      false ->
        simpleArrs = Regex.scan(~r/\[[^\{\}\]\[]+\]/, line)

        case simpleArrs do
          [] ->
            solveSimpleObjects(line)

          matches ->
            tempSums = Enum.map(matches, fn [x] -> Task1.solution(x) end)

            replaced = Regex.replace(~r/\[[^\{\}\]\[]+\]/, line, "_replaced_")

            reduced =
              Enum.reduce(tempSums, replaced, fn x, acc ->
                Regex.replace(~r/_replaced_/, acc, Integer.to_string(x), global: false)
              end)

            # IO.write("array")
            # IO.puts(line)
            # IO.write("reduced")
            # IO.puts(reduced)

            solveSimpleObjects(reduced)
        end
    end
  end
end
