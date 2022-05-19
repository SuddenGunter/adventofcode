defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "abcdffaa" do
    assert Task1.valid("abcdffaa") == true
  end

  test "abidffaa" do
    assert Task1.valid("abidffaa") == false
  end

  test "abcdffba" do
    assert Task1.valid("abcdffba") == false
  end

  test "abcdfffa" do
    assert Task1.valid("abcdfffa") == false
  end

  test "abcdefgh" do
    assert Task1.solution("abcdefgh") == "abcdffaa"
  end

  test "ghijklmn" do
    assert Task1.solution("ghijklmn") == "ghjaabcc"
  end

  test "inc last aaaaaaay" do
    assert Task1.incrementLast("aaaaaaay") == "aaaaaaaz"
  end

  test "inc last aaaaaaaz" do
    assert Task1.incrementLast("aaaaaaaz") == "aaaaaaba"
  end

  test "inc last aaaaaazz" do
    assert Task1.incrementLast("aaaaaazz") == "aaaaabaa"
  end
end
