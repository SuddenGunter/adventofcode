defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.downcase()
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: %{String.t() => %Signal{}}
  def data(contents) do
    lines(contents)
    |> Enum.reduce(%{signals: %{}, line: 0}, fn x, acc ->
      parsed = parseEachLine(x, acc.line)
      %{signals: Map.merge(acc.signals, parsed), line: acc.line + 1}
    end)
    |> Map.get(:signals)
  end

  defp parseEachLine(line, number) do
    [left | [right | _]] = String.split(line, "->", trim: true)
    key = String.trim(right)

    gate =
      cond do
        String.contains?(left, "and") -> :and
        String.contains?(left, "or") -> :or
        String.contains?(left, "lshift") -> :lshift
        String.contains?(left, "rshift") -> :rshift
        String.contains?(left, "not") -> :not
        true -> :value
      end

    # treat as special case and do not create extra noded if value is an integer.
    if gate == :value,
      do: %Signal{
        gate: gate,
        left: parseOperands([left], number)
      }

    operands = String.split(left, Atom.to_string(gate), trim: true)

    operands = parseOperands(operands, number)

    names =
      Map.merge(
        %Signal{
          gate: gate
        },
        operands.names
      )

    Map.merge(
      %{
        key => names
      },
      operands.values
    )
  end

  defp parseOperand(operand, fieldName, line) do
    first = operand |> String.trim()

    case Integer.parse(first) do
      {num, _} ->
        %{
          names: %{fieldName => Atom.to_string(fieldName) <> Integer.to_string(line)},
          values: %{
            (Atom.to_string(fieldName) <> Integer.to_string(line)) => %Signal{
              left: num,
              gate: :value
            }
          }
        }

      :error ->
        %{
          names: %{left: first},
          values: %{}
        }
    end
  end

  defp parseOperands(operands, number) when length(operands) == 1 do
    left = parseOperand(Enum.at(operands, 0), :left, number)

    %{
      names: left.names,
      values: left.values
    }
  end

  defp parseOperands(operands, number) when length(operands) == 2 do
    left = parseOperand(Enum.at(operands, 0), :left, number)
    right = parseOperand(Enum.at(operands, 1), :right, number)

    %{
      names: Map.merge(left.names, right.names),
      values: Map.merge(left.values, right.values)
    }
  end
end
