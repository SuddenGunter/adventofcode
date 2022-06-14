defmodule TaskTest do
  use ExUnit.Case
  doctest Task1

  test "Coment/Dancer 1000s" do
    deers =
      [%Reindeer{speed: 14, runTime: 10, restTime: 127}] ++
        [%Reindeer{speed: 16, runTime: 11, restTime: 162}]

    assert Task2.solution(deers, 1000) == 689
  end
end
