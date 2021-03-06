// source: heyrenee/v1/messages/appointment.proto
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

var heyrenee_v1_messages_provider_pb = require('../../../heyrenee/v1/messages/provider_pb.js');
goog.object.extend(proto, heyrenee_v1_messages_provider_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.heyrenee.v1.messages.Appointment', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.Appointment.ProviderCase', null, global);
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
proto.heyrenee.v1.messages.Appointment = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.heyrenee.v1.messages.Appointment.repeatedFields_, proto.heyrenee.v1.messages.Appointment.oneofGroups_);
};
goog.inherits(proto.heyrenee.v1.messages.Appointment, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.Appointment.displayName = 'proto.heyrenee.v1.messages.Appointment';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.heyrenee.v1.messages.Appointment.repeatedFields_ = [6];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.heyrenee.v1.messages.Appointment.oneofGroups_ = [[2,3]];

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.Appointment.ProviderCase = {
  PROVIDER_NOT_SET: 0,
  PROVIDER_ID: 2,
  PROVIDER_MESSAGE: 3
};

/**
 * @return {proto.heyrenee.v1.messages.Appointment.ProviderCase}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getProviderCase = function() {
  return /** @type {proto.heyrenee.v1.messages.Appointment.ProviderCase} */(jspb.Message.computeOneofCase(this, proto.heyrenee.v1.messages.Appointment.oneofGroups_[0]));
};



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
proto.heyrenee.v1.messages.Appointment.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.Appointment.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.Appointment} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Appointment.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    providerId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    providerMessage: (f = msg.getProviderMessage()) && heyrenee_v1_messages_provider_pb.Provider.toObject(includeInstance, f),
    appointmentId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    date: (f = msg.getDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    addressLinesList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
    city: jspb.Message.getFieldWithDefault(msg, 7, ""),
    state: jspb.Message.getFieldWithDefault(msg, 8, ""),
    zip: jspb.Message.getFieldWithDefault(msg, 9, ""),
    instructions: jspb.Message.getFieldWithDefault(msg, 10, "")
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
 * @return {!proto.heyrenee.v1.messages.Appointment}
 */
proto.heyrenee.v1.messages.Appointment.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.Appointment;
  return proto.heyrenee.v1.messages.Appointment.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.Appointment} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.Appointment}
 */
proto.heyrenee.v1.messages.Appointment.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setProviderId(value);
      break;
    case 3:
      var value = new heyrenee_v1_messages_provider_pb.Provider;
      reader.readMessage(value,heyrenee_v1_messages_provider_pb.Provider.deserializeBinaryFromReader);
      msg.setProviderMessage(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAppointmentId(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDate(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.addAddressLines(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setCity(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setState(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setZip(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setInstructions(value);
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
proto.heyrenee.v1.messages.Appointment.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.Appointment.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.Appointment} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Appointment.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getProviderMessage();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      heyrenee_v1_messages_provider_pb.Provider.serializeBinaryToWriter
    );
  }
  f = message.getAppointmentId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getDate();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getAddressLinesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      6,
      f
    );
  }
  f = message.getCity();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getState();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getZip();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getInstructions();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string provider_id = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getProviderId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setProviderId = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.heyrenee.v1.messages.Appointment.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.clearProviderId = function() {
  return jspb.Message.setOneofField(this, 2, proto.heyrenee.v1.messages.Appointment.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.Appointment.prototype.hasProviderId = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Provider provider_message = 3;
 * @return {?proto.heyrenee.v1.messages.Provider}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getProviderMessage = function() {
  return /** @type{?proto.heyrenee.v1.messages.Provider} */ (
    jspb.Message.getWrapperField(this, heyrenee_v1_messages_provider_pb.Provider, 3));
};


/**
 * @param {?proto.heyrenee.v1.messages.Provider|undefined} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
*/
proto.heyrenee.v1.messages.Appointment.prototype.setProviderMessage = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.heyrenee.v1.messages.Appointment.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.clearProviderMessage = function() {
  return this.setProviderMessage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.Appointment.prototype.hasProviderMessage = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string appointment_id = 4;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getAppointmentId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setAppointmentId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp date = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
*/
proto.heyrenee.v1.messages.Appointment.prototype.setDate = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.clearDate = function() {
  return this.setDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.Appointment.prototype.hasDate = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated string address_lines = 6;
 * @return {!Array<string>}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getAddressLinesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setAddressLinesList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.addAddressLines = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.clearAddressLinesList = function() {
  return this.setAddressLinesList([]);
};


/**
 * optional string city = 7;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getCity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setCity = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string state = 8;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getState = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setState = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string zip = 9;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getZip = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setZip = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string instructions = 10;
 * @return {string}
 */
proto.heyrenee.v1.messages.Appointment.prototype.getInstructions = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Appointment} returns this
 */
proto.heyrenee.v1.messages.Appointment.prototype.setInstructions = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


goog.object.extend(exports, proto.heyrenee.v1.messages);
