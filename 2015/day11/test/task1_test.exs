defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "abcdffaa" do
    assert Task1.valid("abcdffaa") == true
  end

  test "abidffaa" do
    assert Task1.valid("abidffaa") == false
  end
end
