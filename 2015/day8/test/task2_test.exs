defmodule Task2Test do
  use ExUnit.Case
  doctest Task2

  test "empty line" do
    assert Task2.solution(["\"\""]) == 4
  end

  test "abc" do
    assert Task2.solution(["\"abc\""]) == 4
  end

  test "aaaaaa" do
    assert Task2.solution(["\"aaa\\\"aaa\""]) == 6
  end

  test "x27" do
    assert Task2.solution(["\"\\x27\""]) == 5
  end
end
