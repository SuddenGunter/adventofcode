defmodule Day1 do
  def solution([first, second | tail], acc) do
    solution(
      [second | tail],
      if first < second do
        acc + 1
      else
        acc + 0
      end
    )
  end

  def solution([_last], acc) do
    acc
  end

  def solution([], acc) do
    acc
  end
end
