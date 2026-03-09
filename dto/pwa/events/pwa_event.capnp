@0xf03ac6a2556d316d;

using Go = import "/go.capnp";
using Envelope = import "../../common/envelope.capnp".Envelope;
$Go.package("dto");
$Go.import("github.com/ApptroveLabs/apptrove-shared-models/dto/pwa/events");

struct PWAEvent {
  wsid    @0 :Text;
  cuid    @1 :Text;
  mode    @2 :Text;
  created @3 :Int64;
  event   @4 :Event;

  struct Event {
    id        @0  :Text;
    revenue   @1  :Text;
    currency  @2  :Text;
    orderId   @3  :Text;
    productId @4  :Text;
    cCode     @5  :Text;
    param1    @6  :Text;
    param2    @7  :Text;
    param3    @8  :Text;
    param4    @9  :Text;
    param5    @10 :Text;
    param6    @11 :Text;
    param7    @12 :Text;
    param8    @13 :Text;
    param9    @14 :Text;
    param10   @15 :Text;
    ev        @16 :Text;
  }
}

struct PWAEventMessage {
    envelope @0 :Envelope;
    body     @1 :PWAEvent;
}