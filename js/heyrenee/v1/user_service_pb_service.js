// package: heyrenee.v1
// file: heyrenee/v1/user_service.proto

var heyrenee_v1_user_service_pb = require("../../heyrenee/v1/user_service_pb");
var heyrenee_v1_messages_patient_pb = require("../../heyrenee/v1/messages/patient_pb");
var heyrenee_v1_messages_caregiver_pb = require("../../heyrenee/v1/messages/caregiver_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UserService = (function () {
  function UserService() {}
  UserService.serviceName = "heyrenee.v1.UserService";
  return UserService;
}());

UserService.CreatePatient = {
  methodName: "CreatePatient",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.CreatePatientRequest,
  responseType: heyrenee_v1_messages_patient_pb.Patient
};

UserService.GetPatient = {
  methodName: "GetPatient",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.GetPatientRequest,
  responseType: heyrenee_v1_messages_patient_pb.Patient
};

UserService.UpdatePatient = {
  methodName: "UpdatePatient",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.UpdatePatientRequest,
  responseType: heyrenee_v1_messages_patient_pb.Patient
};

UserService.CreateCaregiver = {
  methodName: "CreateCaregiver",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.CreateCaregiverRequest,
  responseType: heyrenee_v1_messages_caregiver_pb.Caregiver
};

UserService.GetCaregiver = {
  methodName: "GetCaregiver",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.GetCaregiverRequest,
  responseType: heyrenee_v1_messages_caregiver_pb.Caregiver
};

UserService.UpdateCaregiver = {
  methodName: "UpdateCaregiver",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_user_service_pb.UpdateCaregiverRequest,
  responseType: heyrenee_v1_messages_caregiver_pb.Caregiver
};

exports.UserService = UserService;

function UserServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserServiceClient.prototype.createPatient = function createPatient(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.CreatePatient, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

UserServiceClient.prototype.getPatient = function getPatient(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.GetPatient, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

UserServiceClient.prototype.updatePatient = function updatePatient(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.UpdatePatient, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

UserServiceClient.prototype.createCaregiver = function createCaregiver(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.CreateCaregiver, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

UserServiceClient.prototype.getCaregiver = function getCaregiver(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.GetCaregiver, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

UserServiceClient.prototype.updateCaregiver = function updateCaregiver(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UserService.UpdateCaregiver, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.UserServiceClient = UserServiceClient;

