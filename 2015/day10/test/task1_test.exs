defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "1" do
    assert Task1.next("1") == "11"
  end

  test "11" do
    assert Task1.next("11") == "21"
  end

  test "21" do
    assert Task1.next("21") == "1211"
  end

  test "1211" do
    assert Task1.next("1211") == "111221"
  end

  test "111221" do
    assert Task1.next("111221") == "312211"
  end
end
