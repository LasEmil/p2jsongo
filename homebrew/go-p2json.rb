require 'rbconfig'
class Go-P2json < Formula
  desc ""
  homepage "https://github.com/LasEmil/go-p2json"
  version "0.0.2"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/LasEmil/go-p2json/releases/download/v0.0.2/properties-to-json_0.0.2_darwin_amd64.zip"
      sha256 "3776707693bfe60fa1e0774fde4320ed43f047024bf2075c7d3ef35844f509fa"
    when /linux/
      url "https://github.com/LasEmil/go-p2json/releases/download/v0.0.2/properties-to-json_0.0.2_linux_amd64.tar.gz"
      sha256 "cb87151cedbbe479f659d271b867edbf48ef9cec5fadddacb88f8dc28d9016e6"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/LasEmil/go-p2json/releases/download/v0.0.2/properties-to-json_0.0.2_darwin_386.zip"
      sha256 "1c7c9214b61c49c71d60d73bc621a5f874f84df2d1aa0dd57ea07c22204acaa3"
    when /linux/
      url "https://github.com/LasEmil/go-p2json/releases/download/v0.0.2/properties-to-json_0.0.2_linux_386.tar.gz"
      sha256 "3d48d95ca3ea1a9996930e0019f6e1e7e921640f3cabb0ec05b4cbb6e54c46be"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "go-p2json"
  end

  test do
    system "go-p2json"
  end

end
