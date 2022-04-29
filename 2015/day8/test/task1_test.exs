defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "empty line" do
    assert Task1.solution(["\"\""]) == 2
  end

  test "abc" do
    assert Task1.solution(["\"abc\""]) == 2
  end

  test "aaaaaa" do
    assert Task1.solution(["\"aaa\\\"aaa\""]) == 3
  end

  test "x27" do
    assert Task1.solution(["\"\\x27\""]) == 5
  end
end
