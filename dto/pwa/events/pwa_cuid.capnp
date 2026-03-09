@0xada1c3e2997cc120;

using Go = import "/go.capnp";
using Envelope = import "../../common/envelope.capnp".Envelope;
$Go.package("dto");
$Go.import("github.com/ApptroveLabs/apptrove-shared-models/dto/pwa/events");

struct PWACuid {
    wsid    @0 :Text;
    cuid    @1 :Text;
    mode    @2 :Text;
    created @3 :Int64;
}

struct PWACuidMessage {
    envelope @0 :Envelope;
    body @1 :PWACuid;
}