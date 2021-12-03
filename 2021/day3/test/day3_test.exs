defmodule Day3Test do
  use ExUnit.Case
  doctest Task1

  test "dinary diagnostic (test case)" do
    assert Task1.solution([
             "00100",
             "11110",
             "10110",
             "10111",
             "10101",
             "01111",
             "00111",
             "11100",
             "10000",
             "11001",
             "00010",
             "01010"
           ]) == 198
  end
end
