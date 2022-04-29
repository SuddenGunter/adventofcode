defmodule Input do
  @spec read(Path.t(), (binary() -> any)) :: {:ok, any()} | {:error, String.t()}
  def read(filename, parse) do
    case File.read(filename) do
      {:ok, contents} -> {:ok, parse.(contents)}
      {:error, reason} -> {:error, "failed to read file '#{filename}': #{reason}"}
    end
  end
end
