@0xc6814fc02d95d55a;

using Go = import "/go.capnp";
using Envelope = import "../common/envelope.capnp".Envelope;
$Go.package("dto");
$Go.import("github.com/ApptroveLabs/apptrove-shared-models/dto/clicks");


struct Click {
  envelope @0 :Envelope;
  clickId  @1 :Text;
  userId   @2 :Text;
  adId     @3 :Text;
  timestamp @4 :Int64;
}