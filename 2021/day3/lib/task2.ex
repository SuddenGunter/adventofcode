use Bitwise

defmodule Task2 do
  def leastPopSelector(oneAtIndexbitToCheck, count) do
    oneAtIndexbitToCheck < count / 2
  end

  def mostPopSelector(oneAtIndexbitToCheck, count) do
    oneAtIndexbitToCheck >= count / 2
  end

  def recursiveSolver(inputs, selector, bitToCheck) do
    if Enum.count(inputs) == 1 do
      Enum.at(inputs, 0)
    else
      oneAtIndexbitToCheck =
        Enum.reduce(
          inputs,
          0,
          fn x, acc ->
            acc + String.to_integer(String.at(x, bitToCheck), 2)
          end
        )

      numbitToCheck =
        if selector.(oneAtIndexbitToCheck, Enum.count(inputs)) do
          1
        else
          0
        end

      continuation =
        Enum.filter(inputs, fn x ->
          String.at(x, bitToCheck) == Integer.to_string(numbitToCheck)
        end)

      recursiveSolver(
        continuation,
        selector,
        bitToCheck + 1
      )
    end
  end

  def solution(inputs) do
    mostPop = recursiveSolver(inputs, &mostPopSelector/2, 0)
    leastPop = recursiveSolver(inputs, &leastPopSelector/2, 0)

    String.to_integer(mostPop, 2) * String.to_integer(leastPop, 2)
  end
end
