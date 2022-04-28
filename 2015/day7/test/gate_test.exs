defmodule GateTest do
  use ExUnit.Case
  doctest GateTest

  test "three vowels is nice" do
    assert Task1.solution(["aaa"]) == 1
  end

  test "duplicates are nice" do
    assert Task1.solution(["ugknbfddgicrmopn"]) == 1
  end

  test "ab cd pq xy are not nice" do
    assert Task1.solution(["haegwjzuvuyypxyu"]) == 0
  end
end
