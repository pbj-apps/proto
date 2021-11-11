// source: heyrenee/v1/messages/refill.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.heyrenee.v1.messages.Refill', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.RefillType', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.heyrenee.v1.messages.Refill = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.messages.Refill, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.Refill.displayName = 'proto.heyrenee.v1.messages.Refill';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.heyrenee.v1.messages.Refill.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.Refill.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.Refill} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Refill.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    medicationId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    prescriptionId: jspb.Message.getFieldWithDefault(msg, 3, ""),
    refillId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    dateTimeRefilled: (f = msg.getDateTimeRefilled()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    pharmacyId: jspb.Message.getFieldWithDefault(msg, 6, ""),
    refillType: jspb.Message.getFieldWithDefault(msg, 7, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.heyrenee.v1.messages.Refill}
 */
proto.heyrenee.v1.messages.Refill.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.Refill;
  return proto.heyrenee.v1.messages.Refill.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.Refill} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.Refill}
 */
proto.heyrenee.v1.messages.Refill.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPatientId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setMedicationId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrescriptionId(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setRefillId(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDateTimeRefilled(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPharmacyId(value);
      break;
    case 7:
      var value = /** @type {!proto.heyrenee.v1.messages.RefillType} */ (reader.readEnum());
      msg.setRefillType(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.heyrenee.v1.messages.Refill.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.Refill.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.Refill} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Refill.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getMedicationId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPrescriptionId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getRefillId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getDateTimeRefilled();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPharmacyId();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getRefillType();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.Refill.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string medication_id = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.Refill.prototype.getMedicationId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setMedicationId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string prescription_id = 3;
 * @return {string}
 */
proto.heyrenee.v1.messages.Refill.prototype.getPrescriptionId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setPrescriptionId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string refill_id = 4;
 * @return {string}
 */
proto.heyrenee.v1.messages.Refill.prototype.getRefillId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setRefillId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp date_time_refilled = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.Refill.prototype.getDateTimeRefilled = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
*/
proto.heyrenee.v1.messages.Refill.prototype.setDateTimeRefilled = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.clearDateTimeRefilled = function() {
  return this.setDateTimeRefilled(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.Refill.prototype.hasDateTimeRefilled = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string pharmacy_id = 6;
 * @return {string}
 */
proto.heyrenee.v1.messages.Refill.prototype.getPharmacyId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setPharmacyId = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional RefillType refill_type = 7;
 * @return {!proto.heyrenee.v1.messages.RefillType}
 */
proto.heyrenee.v1.messages.Refill.prototype.getRefillType = function() {
  return /** @type {!proto.heyrenee.v1.messages.RefillType} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.RefillType} value
 * @return {!proto.heyrenee.v1.messages.Refill} returns this
 */
proto.heyrenee.v1.messages.Refill.prototype.setRefillType = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.RefillType = {
  REFILL_TYPE_UNSPECIFIED: 0,
  REFILL_DELIVERY: 1,
  REFILL_PICKUP: 2
};

goog.object.extend(exports, proto.heyrenee.v1.messages);
