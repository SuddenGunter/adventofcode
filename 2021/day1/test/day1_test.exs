defmodule Day1Test do
  use ExUnit.Case
  doctest Day1

  test "sonar sweep (test case)" do
    assert Day1.solution([199, 200, 208, 210, 200, 207, 240, 269, 260, 263], 0) == 7
  end
end
