@0xe30d60fb3b9d849e;

using Go = import "/go.capnp";
using Envelope = import "../../common/envelope.capnp".Envelope;
$Go.package("dto");
$Go.import("github.com/ApptroveLabs/apptrove-shared-models/dto/pwa/clicks");

struct Header {
  key   @0 :Text;
  value @1 :Text;
}

struct PWAClick {
  wsid    @0 :Text;
  url     @1 :Text;
  headers @2 :List(Header);
  cuid    @3 :Text;
}


struct Click {
  envelope @0 :Envelope;
  body @1 :PWAClick;
}