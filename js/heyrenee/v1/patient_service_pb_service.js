// package: heyrenee.v1
// file: heyrenee/v1/patient_service.proto

var heyrenee_v1_patient_service_pb = require("../../heyrenee/v1/patient_service_pb");
var heyrenee_v1_messages_patient_caregiver_pb = require("../../heyrenee/v1/messages/patient_caregiver_pb");
var heyrenee_v1_messages_patient_health_questionnaire_pb = require("../../heyrenee/v1/messages/patient_health_questionnaire_pb");
var heyrenee_v1_messages_patient_satisfaction_questionnaire_pb = require("../../heyrenee/v1/messages/patient_satisfaction_questionnaire_pb");
var heyrenee_v1_messages_patient_sdoh_questionnaire_pb = require("../../heyrenee/v1/messages/patient_sdoh_questionnaire_pb");
var heyrenee_v1_messages_patient_assessment_pb = require("../../heyrenee/v1/messages/patient_assessment_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PatientService = (function () {
  function PatientService() {}
  PatientService.serviceName = "heyrenee.v1.PatientService";
  return PatientService;
}());

PatientService.CreatePatientCaregiver = {
  methodName: "CreatePatientCaregiver",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.CreatePatientCaregiverRequest,
  responseType: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver
};

PatientService.UpdatePatientCaregiver = {
  methodName: "UpdatePatientCaregiver",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.UpdatePatientCaregiverRequest,
  responseType: heyrenee_v1_messages_patient_caregiver_pb.PatientCaregiver
};

PatientService.ListPatientCaregivers = {
  methodName: "ListPatientCaregivers",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.ListPatientCaregiversRequest,
  responseType: heyrenee_v1_patient_service_pb.ListPatientCaregiversResponse
};

PatientService.CreatePatientAssessment = {
  methodName: "CreatePatientAssessment",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.CreatePatientAssessmentRequest,
  responseType: heyrenee_v1_messages_patient_assessment_pb.PatientAssessment
};

PatientService.ListPatientAssessments = {
  methodName: "ListPatientAssessments",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.ListPatientAssessmentsRequest,
  responseType: heyrenee_v1_patient_service_pb.ListPatientAssessmentsResponse
};

PatientService.CreatePatientSatisfactionQuestionnaire = {
  methodName: "CreatePatientSatisfactionQuestionnaire",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.CreatePatientSatisfactionQuestionnaireRequest,
  responseType: heyrenee_v1_messages_patient_satisfaction_questionnaire_pb.PatientSatisfactionQuestionnaire
};

PatientService.ListPatientSatisfactionQuestionnaires = {
  methodName: "ListPatientSatisfactionQuestionnaires",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesRequest,
  responseType: heyrenee_v1_patient_service_pb.ListPatientSatisfactionQuestionnairesResponse
};

PatientService.CreatePatientHealthQuestionnaire = {
  methodName: "CreatePatientHealthQuestionnaire",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.CreatePatientHealthQuestionnaireRequest,
  responseType: heyrenee_v1_messages_patient_health_questionnaire_pb.PatientHealthQuestionnaire
};

PatientService.ListPatientHealthQuestionnaires = {
  methodName: "ListPatientHealthQuestionnaires",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesRequest,
  responseType: heyrenee_v1_patient_service_pb.ListPatientHealthQuestionnairesResponse
};

PatientService.CreatePatientSdohQuestionnaire = {
  methodName: "CreatePatientSdohQuestionnaire",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.CreatePatientSdohQuestionnaireRequest,
  responseType: heyrenee_v1_messages_patient_sdoh_questionnaire_pb.PatientSdohQuestionnaire
};

PatientService.ListPatientSdohQuestionnaires = {
  methodName: "ListPatientSdohQuestionnaires",
  service: PatientService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesRequest,
  responseType: heyrenee_v1_patient_service_pb.ListPatientSdohQuestionnairesResponse
};

exports.PatientService = PatientService;

function PatientServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PatientServiceClient.prototype.createPatientCaregiver = function createPatientCaregiver(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.CreatePatientCaregiver, {
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

PatientServiceClient.prototype.updatePatientCaregiver = function updatePatientCaregiver(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.UpdatePatientCaregiver, {
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

PatientServiceClient.prototype.listPatientCaregivers = function listPatientCaregivers(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.ListPatientCaregivers, {
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

PatientServiceClient.prototype.createPatientAssessment = function createPatientAssessment(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.CreatePatientAssessment, {
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

PatientServiceClient.prototype.listPatientAssessments = function listPatientAssessments(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.ListPatientAssessments, {
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

PatientServiceClient.prototype.createPatientSatisfactionQuestionnaire = function createPatientSatisfactionQuestionnaire(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.CreatePatientSatisfactionQuestionnaire, {
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

PatientServiceClient.prototype.listPatientSatisfactionQuestionnaires = function listPatientSatisfactionQuestionnaires(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.ListPatientSatisfactionQuestionnaires, {
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

PatientServiceClient.prototype.createPatientHealthQuestionnaire = function createPatientHealthQuestionnaire(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.CreatePatientHealthQuestionnaire, {
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

PatientServiceClient.prototype.listPatientHealthQuestionnaires = function listPatientHealthQuestionnaires(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.ListPatientHealthQuestionnaires, {
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

PatientServiceClient.prototype.createPatientSdohQuestionnaire = function createPatientSdohQuestionnaire(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.CreatePatientSdohQuestionnaire, {
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

PatientServiceClient.prototype.listPatientSdohQuestionnaires = function listPatientSdohQuestionnaires(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PatientService.ListPatientSdohQuestionnaires, {
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

exports.PatientServiceClient = PatientServiceClient;

