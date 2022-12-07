defmodule Fs do
  defmodule File do
    defstruct name: "", size: 0
  end

  defmodule Dir do
    defstruct name: "", size: 0, children: []
  end
end
