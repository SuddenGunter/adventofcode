defmodule Task2Test do
  use ExUnit.Case
  doctest Task2

  test ~s({"d":"red","e":[1,2,3,4],"f":5}) do
    assert Task2.solution(~s({"d":"red","e":[1,2,3,4],"f":5})) == 0
  end

  test ~s([1,2,3]) do
    assert Task2.solution(~s([1,2,3])) == 6
  end

  test ~s([1,{"c":"red","b":2},3]) do
    assert Task2.solution(~s([1,{"c":"red","b":2},3])) == 4
  end

  test ~s([1,5,"red"]) do
    assert Task2.solution(~s([1,5,"red"])) == 6
  end
end
