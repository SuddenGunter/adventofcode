defmodule Task2Test do
  use ExUnit.Case
  doctest Task2

  test "Task2-1" do
    assert Task2.solution(["two1nine"]) == 29
  end

  test "Task2-2" do
    assert Task2.solution(["eightwothree"]) == 83
  end

  test "Task2-3" do
    assert Task2.solution(["zoneight234"]) == 14
  end

  test "Task2-4" do
    assert Task2.solution([
             "two1nine",
             "eightwothree",
             "abcone2threexyz",
             "xtwone3four",
             "4nineeightseven2",
             "zoneight234",
             "7pqrstsixteen"
           ]) == 281
  end
end
