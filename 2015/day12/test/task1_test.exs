defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "[1,2,3]" do
    assert Task1.solution("[1,2,3]") == 6
  end

  test ~s("{"a":2,"b":4}") do
    assert Task1.solution(~s({"a":2,"b":4})) == 6
  end

  test ~s([[[3]]]) do
    assert Task1.solution(~s([[[3]]])) == 3
  end

  test ~s({"a":[-1,1]}) do
    assert Task1.solution(~s({"a":[-1,1]})) == 0
  end

  test ~s({}) do
    assert Task1.solution(~s({})) == 0
  end
end
