defmodule Day2Test do
  use ExUnit.Case
  doctest Day2

  test "sonar sweep (test case)" do
    assert Day2.solution([199, 200, 208, 210, 200, 207, 240, 269, 260, 263], 0, 0) == 5
  end
end
