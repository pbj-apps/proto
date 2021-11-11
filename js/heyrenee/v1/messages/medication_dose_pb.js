// source: heyrenee/v1/messages/medication_dose.proto
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
goog.exportSymbol('proto.heyrenee.v1.messages.MedicationDose', null, global);
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
proto.heyrenee.v1.messages.MedicationDose = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.messages.MedicationDose, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.MedicationDose.displayName = 'proto.heyrenee.v1.messages.MedicationDose';
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
proto.heyrenee.v1.messages.MedicationDose.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.MedicationDose.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.MedicationDose} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.MedicationDose.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    medicationId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    prescriptionId: jspb.Message.getFieldWithDefault(msg, 3, ""),
    medicationDoseId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    timeTaken: (f = msg.getTimeTaken()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
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
 * @return {!proto.heyrenee.v1.messages.MedicationDose}
 */
proto.heyrenee.v1.messages.MedicationDose.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.MedicationDose;
  return proto.heyrenee.v1.messages.MedicationDose.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.MedicationDose} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.MedicationDose}
 */
proto.heyrenee.v1.messages.MedicationDose.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setMedicationDoseId(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTimeTaken(value);
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
proto.heyrenee.v1.messages.MedicationDose.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.MedicationDose.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.MedicationDose} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.MedicationDose.serializeBinaryToWriter = function(message, writer) {
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
  f = message.getMedicationDoseId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getTimeTaken();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string medication_id = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.getMedicationId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.setMedicationId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string prescription_id = 3;
 * @return {string}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.getPrescriptionId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.setPrescriptionId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string medication_dose_id = 4;
 * @return {string}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.getMedicationDoseId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.setMedicationDoseId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp time_taken = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.getTimeTaken = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
*/
proto.heyrenee.v1.messages.MedicationDose.prototype.setTimeTaken = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.MedicationDose} returns this
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.clearTimeTaken = function() {
  return this.setTimeTaken(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.MedicationDose.prototype.hasTimeTaken = function() {
  return jspb.Message.getField(this, 5) != null;
};


goog.object.extend(exports, proto.heyrenee.v1.messages);
