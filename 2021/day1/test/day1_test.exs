defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "Task1: sonar sweep (test case)" do
    assert Task1.solution([199, 200, 208, 210, 200, 207, 240, 269, 260, 263], 0) == 7
  end

  test "Task2: sonar sweep (test case)" do
    assert Task2.solution([199, 200, 208, 210, 200, 207, 240, 269, 260, 263], 0, 0) == 5
  end
end
